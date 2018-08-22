package com.yauritux.repository;

import java.util.Collection;

import org.springframework.data.jpa.repository.JpaRepository;

import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.CardTransaction;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface CardTransactionRepository extends JpaRepository<CardTransaction, Long> {

	Collection<CardTransaction> findByCard(Card card);
}
