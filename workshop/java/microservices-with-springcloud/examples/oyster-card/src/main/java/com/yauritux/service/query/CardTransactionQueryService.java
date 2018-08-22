package com.yauritux.service.query;

import java.util.List;
import java.util.Optional;

import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.CardTransaction;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface CardTransactionQueryService {

	Optional<List<CardTransaction>> findByCard(Card card);
}
