package com.yauritux.service.query.impl;

import static org.junit.Assert.assertTrue;
import static org.mockito.Mockito.when;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.CardTransaction;
import com.yauritux.repository.CardTransactionRepository;
import com.yauritux.service.query.CardTransactionQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
public class CardTransactionQueryServiceImplTest {
	
	private static final String CARD_SN = "1234567890";
	
	@MockBean
	private CardTransactionRepository cardTransactionRepository;
	
	private CardTransactionQueryService cardTransactionQueryService;
	
	@Before
	public void setup() {
		this.cardTransactionQueryService = new CardTransactionQueryServiceImpl(cardTransactionRepository);
	}

	@Test
	public void findByCard_hasSomeTransactions_recordsPresent() {
		Card card = new Card();
		card.setSerialNo(CARD_SN);
		card.setOwner("Yauri Attamimi");
		card.setBalance(BigDecimal.valueOf(500));
		CardTransaction cardTransaction = new CardTransaction();
		cardTransaction.setJourneyType(JourneyType.BUS);
		cardTransaction.setCard(card);
		cardTransaction.setDate(java.sql.Date.valueOf(LocalDateTime.now().toLocalDate()));
		List<CardTransaction> cardTransactions = new ArrayList<>();
		cardTransactions.add(cardTransaction);
		when(cardTransactionRepository.findByCard(card)).thenReturn(cardTransactions);
		assertTrue(cardTransactionQueryService.findByCard(card).isPresent());
	}
}
