package com.yauritux.service.query.impl;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

import java.math.BigDecimal;
import java.util.Optional;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mockito;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.model.entity.Card;
import com.yauritux.repository.CardRepository;
import com.yauritux.service.query.CardQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
public class CardQueryServiceImplTest {
	
	private static final String CARD_SN = "1234567890";
	
	@MockBean
	private CardRepository cardRepository;
	
	private CardQueryService cardQueryService;
	
	@Before
	public void setup() {
		this.cardQueryService = new CardQueryServiceImpl(cardRepository);
	}
	
	@Test
	public void findBySerialNo_existingCard_shouldReturnCard() {
		Card card = new Card();
		card.setSerialNo(CARD_SN);
		card.setBalance(BigDecimal.valueOf(150));
		Mockito.when(cardRepository.findBySerialNo(CARD_SN))
			.thenReturn(card);
		
		Optional<Card> found = cardQueryService.findBySerialNo(CARD_SN);
		assertTrue(found.isPresent());
		assertThat(found.get().getSerialNo()).isEqualTo(card.getSerialNo());
	}
	
	@Test
	public void findBySerialNo_nonExistingCard_cardNotFound() {
		Mockito.when(cardRepository.findBySerialNo(CARD_SN)).thenReturn(null);
		
		Optional<Card> notFound = cardQueryService.findBySerialNo(CARD_SN);
		assertFalse(notFound.isPresent());
	}
}
